"use client";

import { Button } from "@/components/ui/button";
import {
	DropdownMenu,
	DropdownMenuContent,
	DropdownMenuRadioGroup,
	DropdownMenuRadioItem,
	DropdownMenuSeparator,
	DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import Link from "next/link";
import useMe from "@/hooks/useMe";

export function AccountSettingsDropdown() {
	const { name } = useMe();

	return (
		<DropdownMenu>
			<DropdownMenuTrigger asChild>
				<Button variant="outline">{name}</Button>
			</DropdownMenuTrigger>
			<DropdownMenuContent className="w-56">
				<DropdownMenuSeparator />
				<DropdownMenuRadioGroup>
					<DropdownMenuRadioItem value="settings">
						<Link href="/settings">Settings</Link>
					</DropdownMenuRadioItem>
					<DropdownMenuRadioItem value="logout">
						<Link href="/logout">Logout</Link>
					</DropdownMenuRadioItem>
				</DropdownMenuRadioGroup>
			</DropdownMenuContent>
		</DropdownMenu>
	);
}

export default AccountSettingsDropdown;
