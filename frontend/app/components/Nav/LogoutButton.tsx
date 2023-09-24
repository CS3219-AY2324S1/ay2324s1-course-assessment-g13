import { Button } from "@nextui-org/button";
import Link from "next/link";

export default function LogoutButton({handleLogout}) {
    return (
        <Button  
            variant="bordered" 
            color="default"
            onClick={handleLogout}
        >
            Logout
        </Button>
    )
}