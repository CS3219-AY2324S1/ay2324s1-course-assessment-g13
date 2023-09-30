import SignupModal from './components/modal/signupModal';
import LoginModal from "./components/modal/loginModal";

export const metadata = {
  title: 'PeerPrep',
};

export default function Home() {
  return (
    <div className="flex flex-col justify-center items-center text-center my-12">
      <p className="text-6xl font-extrabold my-6">Elevate Your Tech Interview with PeerPrep!</p>
      <p className="text-2xl font-semibold my-6">
        Excel in your technical interviews through collaborative mock interviews and question
        tracking on PeerPrep.
      </p>

      {/*<SignupModal />*/}
        <LoginModal/>
    </div>
  );
}
