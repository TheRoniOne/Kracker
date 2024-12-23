import { JSONFetcher } from "@/lib/fetcher";
import type { Me } from "@/types/me";
import useSWR from "swr";

export default function useMe(): Me {
	const { data } = useSWR("me", JSONFetcher, {
		refreshInterval: 1000 * 60 * 5,
	});

	return {
		name: data?.name,
		permissions: data?.permissions,
	};
}
