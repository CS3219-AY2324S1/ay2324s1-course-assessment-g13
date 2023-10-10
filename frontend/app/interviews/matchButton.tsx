import {Button, user} from "@nextui-org/react";
import {notifyError, notifySuccess, notifyWarning} from "../components/toast/notifications";
import { useEffect, useState } from "react";
import {useSelector} from "react-redux";
import {selectPreferenceState} from "../libs/redux/slices/matchPreferenceSlice";
import {POST} from "../libs/axios/axios";
import {selectUsername} from "../libs/redux/slices/userSlice";

export default function MatchButton({inQueue, setInQueue}) {
  const [seconds, setSeconds] = useState(0);
  const timeLimit = 30;
  const preferenceState = useSelector(selectPreferenceState)
  const userState = useSelector(selectUsername);

  console.log(`User: ${userState} | User preference: ${preferenceState}`)
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

  const startQueue = () => {
    setInQueue(true);
    getMatch().then(r => {
      if (r.status == 200) {
        const payload = r.data;
        // If there is a match, notify success and redirect
        if (payload["match_user"] != "") {
          notifySuccess(`Matched with ${payload["match_user"]}`);
          setInQueue(false);
          setSeconds(0);
          // TODO perform redirection here based on payload redirect url
        } else {
          matchNotfound()
        }
      } else {
        // If failed to call matching service, cancel queue
        cancelQueue();
        return
      }
    });
  }

  const matchNotfound = () => {
    setInQueue(false);
    setSeconds(0);
    notifyError("Timeout! No suitable partner was found")
  };

  const getMatch = async () => {
    return await POST("match", {
      "username":`${userState}`,
      "match_criteria":`${preferenceState.toLowerCase()}`
    });
  };

  const cancelQueue = () => {
    setInQueue(false);
    setSeconds(0);
    notifyWarning("Queue cancelled!");
  }

  return  (
    <div className="mt-6 mx-auto w-2/5 flex justify-between">
      <Button 
        color="primary" 
        variant="solid" 
        className="text-lg py-5 mx-2" 
        fullWidth={true}
        onPress={startQueue}
        isLoading={inQueue}
      >
        {inQueue ? 'Matching' : 'Match'}
      </Button>
      {inQueue && 
      <Button
        color="danger"
        variant="solid"
        className="text-lg py-5 mx-2"
        fullWidth={true}
        onPress={cancelQueue}
      >
        Cancel
      </Button>}
    </div>
  );
}
