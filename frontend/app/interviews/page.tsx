'use client';
import { useEffect, useState } from "react";
import MatchButton from "./matchButton";
import SetPreferencesModal from "./preferenceModal";
import { notifySuccess } from "../components/toast/notifications";

export default function Interviews() {
  const [inQueue, setInQueue] = useState(false);
  const [isMatch, setIsMatch] = useState(false);

  // TODO: When match found set this to true
  useEffect(() => {
    if (isMatch) {
      setInQueue(false);
      setIsMatch(false);
      notifySuccess("Match Found, Redirecting to Collaboration Room...");
    }
  }, [isMatch]);
  
  return (
    <>
      <div className="mx-auto px-6 max-w-7xl flex flex-col justify-center text-center my-12">
        <div className="flex justify-between items-center mb-5">
          <span className="text-3xl">Interviews</span>
          <SetPreferencesModal />
        </div>
        {!inQueue && (
          <div>
            <p className="text-3xl font-extrabold my-6">Looking for a practice partner to conduct a mock interview?</p>
            <p className="text-2xl font-semibold my-6">
              Customize your preferences using the preferences button, and hit match to get started!
            </p>
          </div>
        )}
        {inQueue && (
          <div>
            <p className="text-3xl font-extrabold my-6">Looking for a partner...</p>
            <p className="text-2xl font-semibold my-6">
              You can end the search by clicking the cancel button!
            </p>
          </div>
        )}
        <MatchButton inQueue={inQueue} setInQueue={setInQueue}/>
      </div>
    </>
  );
}
