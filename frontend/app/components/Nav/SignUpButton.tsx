import { Button } from "@nextui-org/button";
import Link from "next/link";

export default function SignUpButton({width=null}) {
    return (
        <Button  
            color="primary"
            as={Link}
            href="/signup"
            className={width}
        >
            Sign Up
        </Button>
    )
}
