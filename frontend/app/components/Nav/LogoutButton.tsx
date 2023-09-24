import { Button } from "@nextui-org/button";
import Link from "next/link";

export default function LogoutButton({handleLogout}) {
    return (
        <Button  
            variant="light" 
            color="default"
            onClick={handleLogout}
        >
            Logout
        </Button>
    )
}