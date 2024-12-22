import qs from "qs";

const API_BASE = process.env.NEXT_PUBLIC_BACKEND;

const buildURL = (endpoint: string): string => {
	return API_BASE + endpoint;
};

const fetcher = async (url: string | URL, options: RequestInit = {}) => {
	return fetch(url, {
		...options,
	}).then((r) => {
		if (r.status === 401) {
			console.log("Response status was 401 should redirect to signin");
			window.location.href = "/login";
		}

		return r;
	});
};

export type Params =
	| Record<
			string,
			string | string[] | number | number[] | undefined | Date | boolean | null
	  >
	| undefined;

const encodeURL = (url: string, queryParams: Params = undefined) => {
	if (queryParams)
		return `${url}?${qs.stringify(queryParams, { arrayFormat: "comma" })}`;

	return url;
};

const fe = {
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

export { encodeURL, fe, fetcher, buildURL };
