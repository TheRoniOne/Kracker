import AccountSettingsDropdown from "@/components/AccountSettingsDropDown";
import { ModeToggle } from "./ModeToggle";

export default function NavBar() {
	return (
		<nav className="flex flex-row justify-between m-3">
			<div>Projects</div>
			<div className="flex flex-row gap-2">
				<AccountSettingsDropdown />
				<ModeToggle />
			</div>
		</nav>
	);
}
