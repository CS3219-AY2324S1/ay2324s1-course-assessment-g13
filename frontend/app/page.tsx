import { Button } from '@nextui-org/button';
import { Link } from '@nextui-org/link';

export const metadata = {
  title: 'PeerPrep',
};

export default function Home() {

  return (
    <div className="flex flex-col justify-center items-center text-center mt-40 max-w-4xl mx-auto">
      <p className="text-6xl font-extrabold my-6">Elevate Your Tech Interview with PeerPrep!</p>
      <p className="text-2xl font-semibold my-6">
        Excel in your technical interviews through collaborative mock interviews and question
        tracking on PeerPrep.
      </p>

      <Button  
            color="primary"
            as={Link}
            href="/signup"
            className={"w-2/5"}
        >
            Sign Up
        </Button>
    </div>
  );
}
