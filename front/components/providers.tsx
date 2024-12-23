"use client";

import { Provider } from "jotai";
import type { ReactNode, ComponentProps } from "react";
import { ThemeProvider as NextThemesProvider } from "next-themes";

export function JotaiProvider({
	children,
}: Readonly<{
	children: ReactNode;
}>) {
	return <Provider>{children}</Provider>;
}

export function ThemeProvider({
	children,
	...props
}: Readonly<ComponentProps<typeof NextThemesProvider>>) {
	return <NextThemesProvider {...props}>{children}</NextThemesProvider>;
}
