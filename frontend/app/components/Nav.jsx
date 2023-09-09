import { Link } from "@nextui-org/link";
import { Button } from "@nextui-org/button";

import { 
    Navbar,
    NavbarBrand,
    NavbarContent,
    NavbarItem,
} from "@nextui-org/navbar";

export default function Nav() {
    return (
        <Navbar
            isBordered
            maxWidth="xl"
            height="10vh"
            classNames={{
            item: [
                "flex",
                "relative",
                "h-full",
                "items-center",
                "data-[active=true]:after:content-['']",
                "data-[active=true]:after:absolute",
                "data-[active=true]:after:bottom-0",
                "data-[active=true]:after:left-0",
                "data-[active=true]:after:right-0",
                "data-[active=true]:after:h-[3px]",
                "data-[active=true]:after:rounded-[2px]",
                "data-[active=true]:after:bg-primary",
            ],
            }}
        >
            <NavbarBrand>
                <p className="font-bold text-inherit">PeerPrep</p>
            </NavbarBrand>
            <NavbarContent className="hidden sm:flex gap-4" justify="center">
                <NavbarItem isActive>
                    <Link href="#" aria-current="page">
                        Questions
                    </Link>
                </NavbarItem>
                <NavbarItem>
                    <Link color="foreground" href="#">
                        Interviews
                    </Link>
                </NavbarItem>
            </NavbarContent>
            <NavbarContent justify="end">
                <NavbarItem className="hidden lg:flex">
                    <Link href="#">Login</Link>
                </NavbarItem>
                <NavbarItem>
                    <Button as={Link} color="primary" href="#" variant="flat">
                        Sign Up
                    </Button>
                </NavbarItem>
            </NavbarContent>
        </Navbar>
    );
}