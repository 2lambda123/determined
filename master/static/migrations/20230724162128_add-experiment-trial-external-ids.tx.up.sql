ALTER TABLE experiments
    ADD COLUMN external_experiment_id TEXT UNIQUE NULL;
ALTER TABLE trials
    ADD COLUMN external_trial_id TEXT UNIQUE NULL;
