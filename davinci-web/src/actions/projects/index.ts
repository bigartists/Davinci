import type { IProject } from 'src/types/project/type';

import useSWR from 'swr';
import { useMemo } from 'react';

import { fetcher } from 'src/utils/axios';

export function useGetProjects() {
  const URL = '/projects';
  const { data, isLoading, error, isValidating } = useSWR(URL, fetcher, {
    refreshInterval: 100000,
  });

  const memoizedValue = useMemo(
    () => ({
      projects: (data?.payload as IProject[]) || [],
      projectsLoading: isLoading,
      projectsError: error,
      projectsValidating: isValidating,
      projectsEmpty: !isLoading && !data?.payload?.length,
    }),
    [data?.payload, error, isLoading, isValidating]
  );

  return memoizedValue;
}
