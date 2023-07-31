CREATE OR REPLACE FUNCTION autoupdate_exp_best_trial_metrics() RETURNS trigger AS $$
BEGIN
    WITH bt AS (SELECT id, best_validation_id FROM trials WHERE experiment_id = NEW.experiment_id ORDER BY searcher_metric_value_signed LIMIT 1)
    UPDATE experiments SET best_trial_id = bt.id, 
    validation_metrics = 
    (
        WITH metrics AS (
            SELECT summary_metrics->'validation_metrics' AS jsonb, searcher_metric_value_signed=searcher_metric_value AS sign FROM trials WHERE id = 41575
        ), 
        metrics_values AS (
            SELECT jsonb_object_keys(jsonb) AS metrics_key, sign, jsonb -> jsonb_object_keys(jsonb) ->> 'min' AS min, jsonb -> jsonb_object_keys(jsonb) ->> 'max' AS max, jsonb -> jsonb_object_keys(jsonb) ->> 'type' AS type FROM metrics),
        result AS (
        SELECT metrics_key, CASE sign when true then min else max end AS metrics_value, type FROM metrics_values WHERE type = 'number'
        ) SELECT json_object_agg(metrics_key, metrics_value::float8) FROM result
    )
    WHERE experiments.id = NEW.experiment_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;



CREATE TABLE public.exp_metrics_name (
    id SERIAL PRIMARY KEY,
    project_id INT REFERENCES projects(id) ON DELETE CASCADE NOT NULL,
    experiment_id INT REFERENCES experiments(id) ON DELETE CASCADE NOT NULL,
    vname JSON
);

CREATE INDEX ix_metrics_name_project_id ON exp_metrics_name USING btree (project_id);
CREATE UNIQUE INDEX ix_metrics_name_experiment_id_unique ON exp_metrics_name(experiment_id);

INSERT INTO public.exp_metrics_name (project_id, experiment_id, vname) (
    WITH validation_metrics_names AS (
        SELECT array_to_json(array_agg(DISTINCT names)) AS name, e.id AS experiment_id
        FROM trials t, experiments e, raw_validations v,
            LATERAL jsonb_object_keys(v.metrics->'validation_metrics') AS names
        WHERE t.best_validation_id=v.id AND e.id = t.experiment_id 
        GROUP BY e.id)
    SELECT   
        e.project_id AS project_id, 
        e.id AS experiment_id, 
        COALESCE(validation_metrics_names.name, '[]'::json) AS vname 
    FROM experiments e, validation_metrics_names
    WHERE 
        validation_metrics_names.experiment_id = e.id
);