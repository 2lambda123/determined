import { useEffect, useState } from 'react';

import { Loadable, Loaded, NotLoaded } from 'utils/loadable';

export const useAsync = <T>(
  cb: () => Promise<T | typeof NotLoaded>,
  deps: React.DependencyList,
): Loadable<T> => {
  const [retVal, setRetVal] = useState<Loadable<T>>(NotLoaded);

  useEffect(() => {
    let mounted = true;
    setRetVal(NotLoaded);
    (async () => {
      const newVal = await cb();
      if (typeof newVal === 'object' && newVal !== null && '_tag' in newVal) {
        return;
      }
      if (mounted) {
        setRetVal(Loaded(newVal));
      }
    })();
    return () => {
      mounted = false;
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [deps]);

  return retVal;
};

export default useAsync;
