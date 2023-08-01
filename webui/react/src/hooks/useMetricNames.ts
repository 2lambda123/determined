import { useEffect, useState } from 'react';

import { XAxisDomain } from 'components/kit/LineChart/XAxisFilter';
import { V1ExpMetricNamesResponse } from 'services/api-ts-sdk';
import { detApi } from 'services/apiConfig';
import { readStream } from 'services/utils';
import { Metric, MetricType, OldMetric } from 'types';
import { isEqual } from 'utils/data';
import { Loadable, Loaded, NotLoaded } from 'utils/loadable';
import { metricKeyToMetric, metricToKey } from 'utils/metric';
import { metricSorter } from 'utils/metric';

export const asOldMetrics = (metrics: Metric[]): OldMetric[] => {
  return metrics.filter((metric) => {
    return ([MetricType.Training, MetricType.Validation] as string[]).includes(metric.type);
  }) as OldMetric[];
};

export const asLoadableOldMetrics = (loadable: Loadable<Metric[]>): Loadable<OldMetric[]> => {
  return Loadable.map(loadable, (metrics) => asOldMetrics(metrics));
};

const useMetricNames = (
  experimentIds: number[],
  errorHandler?: (e: unknown) => void,
): Loadable<Metric[]> => {
  const [metrics, setMetrics] = useState<Loadable<Metric[]>>(NotLoaded);
  const [actualExpIds, setActualExpIds] = useState<number[]>([]);

  useEffect(
    () => setActualExpIds((prev) => (isEqual(prev, experimentIds) ? prev : experimentIds)),
    [experimentIds],
  );

  useEffect(() => {
    if (actualExpIds.length === 0) {
      setMetrics(Loaded([]));
      return;
    }
    const canceler = new AbortController();

    // We do not want to plot any x-axis metric values as y-axis data
    const xAxisMetrics = Object.values(XAxisDomain).map((v) => v.toLowerCase());

    readStream<V1ExpMetricNamesResponse>(
      detApi.StreamingInternal.expMetricNames(actualExpIds, undefined, {
        signal: canceler.signal,
      }),
      (event: V1ExpMetricNamesResponse) => {
        if (!event) return;
        /**
         * The metrics endpoint can intermittently send empty lists,
         * so we keep track of what we have seen on our end and
         * only add new metrics we have not seen to the list.
         */
        const newMetrics = (event.metricNames ?? [])
          .filter((metric) => !xAxisMetrics.includes(metric.name))
          .map((metric) => ({ name: metric.name, type: metric.group } as Metric));

        if (newMetrics.length > 0) {
          setMetrics((prevMetrics) => {
            /**
             * Since we may intermittently receive a subset of all available
             * metrics or an empty list of metrics we must merge the new and
             * previous metrics to accurately determine if any new metrics have
             * not been seen before.
             */
            const previousMetrics = Loadable.getOrElse([], prevMetrics);
            const previousMetricsSet = Loadable.getOrElse([], prevMetrics).reduce(
              (acc, cur) => acc.add(metricToKey(cur)),
              new Set<string>(),
            );
            const updatedMetricsSet = [...newMetrics, ...previousMetrics].reduce(
              (acc, cur) => acc.add(metricToKey(cur)),
              new Set<string>(),
            );

            if (isEqual(previousMetricsSet, updatedMetricsSet)) return prevMetrics;

            return Loaded(
              Array.from(updatedMetricsSet)
                .map((metricKey) => metricKeyToMetric(metricKey))
                .sort(metricSorter),
            );
          });
        }
      },
      errorHandler,
    );
    return () => canceler.abort();
  }, [actualExpIds, errorHandler]);

  return metrics;
};

export default useMetricNames;
