import qs from "qs";

const API_BASE = process.env.NEXT_PUBLIC_BACKEND;

function buildURL(endpoint: string): string {
	return API_BASE + endpoint;
}

export async function fetcher(url: string | URL, options: RequestInit = {}) {
	return fetch(url, {
		...options,
	}).then((r) => {
		if (r.status === 401) {
			console.log("Response status was 401 should redirect to signin");
			window.location.href = "/login";
		}

		return r;
	});
}

export async function JSONFetcher(url: string, options: RequestInit = {}) {
	const r = await fetcher(url, options);
	return await r.json();
}

type Params =
	| Record<
			string,
			string | string[] | number | number[] | undefined | Date | boolean | null
	  >
	| undefined;

function encodeURL(url: string, queryParams: Params = undefined) {
	if (queryParams)
		return `${url}?${qs.stringify(queryParams, { arrayFormat: "comma" })}`;

	return url;
}

export const fe = {
	get: async (url: string, queryParams: Params = undefined) => {
		let newUrl = buildURL(url);
		newUrl = encodeURL(newUrl, queryParams);

		return fetcher(newUrl, { method: "GET" });
	},
	post: async (url: string, body = {}) => {
		return fetcher(buildURL(url), {
			method: "POST",
			body: JSON.stringify(body),
		});
	},
	put: async (url: string, body = {}) => {
		return fetcher(buildURL(url), {
			method: "PUT",
			body: JSON.stringify(body),
		});
	},
	patch: async (url: string, body = {}) => {
		return fetcher(buildURL(url), {
			method: "PATCH",
			body: JSON.stringify(body),
		});
	},
	del: async (url: string, body = {}) => {
		return fetcher(buildURL(url), {
			method: "DELETE",
			body: JSON.stringify(body),
		});
	},
};
