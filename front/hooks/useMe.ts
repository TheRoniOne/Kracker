import { JSONFetcher } from "@/lib/fetcher";
import type { Me } from "@/types/me";
import useSWR from "swr";

export default function useMe(): Me {
	const { data } = useSWR("me", JSONFetcher, { refreshInterval: 5000 });

	return {
		name: data?.name,
		permissions: data?.permissions,
	};
}
