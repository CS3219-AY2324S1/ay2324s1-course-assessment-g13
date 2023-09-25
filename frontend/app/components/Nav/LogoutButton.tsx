import { Button } from "@nextui-org/button";

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