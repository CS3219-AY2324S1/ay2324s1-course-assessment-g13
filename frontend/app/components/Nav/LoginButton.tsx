import { Button } from "@nextui-org/button";
import Link from "next/link";

export default function LoginButton() {
    return (
        <Button  
            variant="bordered" 
            color="default"
            as={Link}
            href="/auth/login"
            className="hover:red"
        >
            Login
        </Button>
    )
}