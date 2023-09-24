import { Button } from "@nextui-org/button";
import Link from "next/link";

export default function LoginButton() {
    return (
        <Button  
            variant="light" 
            color="default"
            as={Link}
            href="/auth/login"
        >
            Login
        </Button>
    )
}