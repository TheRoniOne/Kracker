import AccountSettingsDropdown from "@/components/AccountSettingsDropDown";

export default function NavBar() {
	return (
		<nav className="flex flex-row justify-between m-3">
			<div>Projects</div>
			<AccountSettingsDropdown />
		</nav>
	);
}
