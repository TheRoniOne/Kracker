import Link from "next/link";
import { Button } from "@/components/ui/button";

export default function Home() {
	return (
		<div className="h-screen grid place-content-center text-4xl gap-y-10">
			<h1>Welcome to Kracker</h1>
			<div className="flex place-content-evenly">
				<Button asChild className="text-3xl p-4">
					<Link href="/register">Sign Up</Link>
				</Button>
				<Button asChild className="text-3xl p-4">
					<Link href="/login">Log In</Link>
				</Button>
			</div>
		</div>
	);
}
