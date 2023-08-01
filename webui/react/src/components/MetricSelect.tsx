import { RefSelectProps } from 'antd/es/select';
import React, { useCallback, useMemo, useRef, useState } from 'react';

import Select, { OptGroup, Option, SelectValue } from 'components/kit/Select';
import { Metric } from 'types';
import {
  getMetricName,
  metricKeyToMetric,
  metricSorter,
  metricToKey,
  metricWithTypeToKey,
} from 'utils/metric';

import BadgeTag from './BadgeTag';

const allOptionId = 'ALL_RESULTS';
const resetOptionId = 'RESET_RESULTS';

type SingleHandler = (value: Metric) => void;
type MultipleHandler = (value: Metric[]) => void;

interface Props {
  defaultMetrics: Metric[];
  label?: string;
  metrics: Metric[];
  multiple?: boolean;
  onChange?: SingleHandler | MultipleHandler;
  value?: Metric | Metric[];
  width?: number;
}

const filterFn = (search: string, metricName: string) => {
  return metricName.toLocaleLowerCase().indexOf(search.toLocaleLowerCase()) !== -1;
};

const MetricSelect: React.FC<Props> = ({
  defaultMetrics,
  label = 'Metrics',
  metrics,
  multiple,
  value,
  width = 200,
  onChange,
}: Props) => {
  const [filterString, setFilterString] = useState('');
  const selectRef = useRef<RefSelectProps>(null);

  const metricsByType = useMemo(() => {
    const groups = metrics.reduce((acc, metric) => {
      acc[metric.type] = acc[metric.type] || [];
      acc[metric.type].push(metric.name);
      return acc;
    }, {} as Record<string, string[]>);
    return Object.keys(groups).map((key) => {
      return { metrics: groups[key], type: key };
    });
  }, [metrics]);

  const metricValues = useMemo(() => {
    if (multiple && Array.isArray(value)) return value.map((metric) => metricToKey(metric));
    if (!multiple && !Array.isArray(value) && value) return metricToKey(value);
    return undefined;
  }, [multiple, value]);

  const totalNumMetrics = useMemo(() => {
    return metrics.length;
  }, [metrics]);

  /**
   * visibleMetrics should always match the list of metrics that antd displays to
   * the user, including any filtering.
   */
  const visibleMetrics = useMemo(() => {
    return metrics.filter((metric: Metric) => {
      return filterFn(filterString, metric.name);
    });
  }, [metrics, filterString]);

  const handleMetricSelect = useCallback(
    (newValue: SelectValue) => {
      if (!onChange) return;

      if ((newValue as string) === allOptionId) {
        (onChange as MultipleHandler)(visibleMetrics.sort(metricSorter));
        selectRef.current?.blur();
        return;
      }
      if ((newValue as string) === resetOptionId) {
        (onChange as MultipleHandler)(defaultMetrics.sort(metricSorter));
        selectRef.current?.blur();
        return;
      }

      const metric = metricKeyToMetric(newValue as string);
      if (multiple) {
        const newMetric = Array.isArray(value) ? [...value] : [];
        if (newMetric.indexOf(metric) === -1) newMetric.push(metric);
        (onChange as MultipleHandler)(newMetric.sort(metricSorter));
      } else {
        (onChange as SingleHandler)(metric);
      }
    },
    [multiple, onChange, value, visibleMetrics, defaultMetrics],
  );

  const handleMetricDeselect = useCallback(
    (newValue: SelectValue) => {
      if (!onChange || !multiple) return;
      if (!Array.isArray(value) || value.length <= 1) return;

      const newMetric = Array.isArray(value) ? [...value] : [];
      const index = newMetric.findIndex((metric) => metricToKey(metric) === newValue);
      if (index !== -1) newMetric.splice(index, 1);
      (onChange as MultipleHandler)(newMetric.sort(metricSorter));
    },
    [multiple, onChange, value],
  );

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const handleFiltering = useCallback((search: string, option: any) => {
    if (option.key === allOptionId || option.key === resetOptionId) return true;
    if (!option.value) return false;

    const metric = metricKeyToMetric(option.value);
    if (metric === undefined) return false;

    return filterFn(search, metric.name);
  }, []);

  const handleSearchInputChange = (searchInput: string) => setFilterString(searchInput);

  const handleBlur = () => setFilterString('');

  const allOption = useMemo(() => {
    const numVisibleOptions = visibleMetrics.length;
    const allOptionLabel =
      numVisibleOptions === totalNumMetrics ? 'All' : `All ${numVisibleOptions} results`;
    return (
      <Option key={allOptionId} value={allOptionId}>
        <BadgeTag label={allOptionLabel} />
      </Option>
    );
  }, [totalNumMetrics, visibleMetrics]);

  return (
    <Select
      disableTags
      filterOption={handleFiltering}
      label={label}
      mode={multiple ? 'multiple' : undefined}
      ref={selectRef}
      value={metricValues}
      width={width}
      onBlur={handleBlur}
      onDeselect={handleMetricDeselect}
      onSearch={handleSearchInputChange}
      onSelect={handleMetricSelect}>
      {multiple && visibleMetrics.length > 0 && (
        <Option key={resetOptionId} value={resetOptionId}>
          <BadgeTag label="Reset to Default" />
        </Option>
      )}
      {multiple && visibleMetrics.length > 1 && allOption}
      {metricsByType.map((group) => (
        <OptGroup key={group.type} label={group.type}>
          {group.metrics.map((metric) => {
            const value = metricWithTypeToKey(metric, group.type);
            return (
              <Option key={value} value={value}>
                <BadgeTag label={getMetricName(metric)} tooltip={group.type}>
                  {group.type.substring(0, 1).toUpperCase()}
                </BadgeTag>
              </Option>
            );
          })}
        </OptGroup>
      ))}
    </Select>
  );
};

export default MetricSelect;
