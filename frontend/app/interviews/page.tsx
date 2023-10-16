'use client';
import { useEffect, useState } from "react";
import MatchButton from "./matchButton";
import SetPreferencesModal from "./preferenceModal";
import {notifyError, notifySuccess} from "../components/toast/notifications";
import {useSelector} from "react-redux";
import {selectPreferenceState} from "../libs/redux/slices/matchPreferenceSlice";

export default function Interviews() {
  const [inQueue, setInQueue] = useState(false);
  const [isCancelled, setIsCancelled] = useState(false);
  const [shouldNotifyCancelled, setShouldNotifyCancelled] = useState(false);
  const [seconds, setSeconds] = useState(0);
  const userPreference = useSelector(selectPreferenceState);
  const timeLimit = 30;

  const matchNotfound = () => {
    setSeconds(0);
    setInQueue(false);
    setShouldNotifyCancelled(true);
  };

  useEffect(() => {
    if (!isCancelled && shouldNotifyCancelled) {
      notifyError("Timeout! No suitable partner was found");
    }
  }, [isCancelled, shouldNotifyCancelled]);

  useEffect(() => {
    if (seconds == timeLimit) {
      matchNotfound();
    }
    let timer = null;
    if (inQueue) {
      timer = setInterval(() => {
        setSeconds((seconds) => seconds + 1);
      }, 1000);
    }
    return () => {
      clearInterval(timer);
    };
  }, [inQueue, seconds]);

  return (
    <>
      <div className="mx-auto px-6 max-w-7xl flex flex-col justify-center text-center my-12">
        <div className="flex justify-between items-center mb-5">
          <span className="text-3xl">Interviews</span>
          <span className="text-2xl">Current Preference: {userPreference}</span>
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
            <p className="text-2xl font-semibold my-6 text-teal-400">Time Elapsed: {seconds} seconds</p>
            <p className="text-2xl font-semibold my-6">
              You can end the search by clicking the cancel button!
            </p>
          </div>
        )}
        <MatchButton inQueue={inQueue} setInQueue={setInQueue} setSeconds={setSeconds}
                     matchNotfound={matchNotfound} setIsCancelled={setIsCancelled}
                    setShouldNotifyCancelled={setShouldNotifyCancelled} />
      </div>
    </>
  );
}
