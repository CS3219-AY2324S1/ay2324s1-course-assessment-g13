import {Button} from "@nextui-org/react";
import {notifySuccess, notifyWarning} from "../components/toast/notifications";
import {useSelector} from "react-redux";
import {selectPreferenceState} from "../libs/redux/slices/matchPreferenceSlice";
import {POST} from "../libs/axios/axios";
import {selectUsername} from "../libs/redux/slices/userSlice";
import { useRouter } from 'next/navigation';

export default function MatchButton({inQueue, setInQueue, setSeconds, matchNotfound, setIsCancelled, setShouldNotifyCancelled, active, setActive}) {
  const preferenceState = useSelector(selectPreferenceState)
  const userState = useSelector(selectUsername);
  const router = useRouter();

  const startQueue = () => {
    setInQueue(true);
    setIsCancelled(false);
    setShouldNotifyCancelled(false);
    getMatch().then(r => {
      if (r.status == 200) {
        const payload = r.data;
        // If there is a match, notify success and redirect
        if (payload["match_user"] != "") {
          setActive(true)
          notifySuccess(`Matched with ${payload["match_user"]}, redirecting to collaboration room...`);
          setInQueue(false);
          setSeconds(0);
          redirectToCollab(payload["room_id"]);
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

  const redirectToCollab = (room_id : string) => {
    // Just to match with toast timer, preference whether to instant redirect or not
    const redirectTimer = setTimeout(() => {
      router.push(`/collab/?room_id=${room_id}`);
    }, 1000);

    return () => {
      clearTimeout(redirectTimer);
    }
  }

  const getMatch = async () => {
    return await POST("/match/find", {
      "username":`${userState}`,
      "match_criteria":`${preferenceState.toLowerCase()}`
    });
  };

  const cancelMatch = async () => {
    return await POST("/match/cancel", {
      "username": `${userState}`,
      "match_criteria": `${preferenceState.toLowerCase()}`
    });
  }

  const cancelQueue = () => {
    // Sends API request to matching service to indicate user has cancelled
    cancelMatch().then(res => {
        setIsCancelled(true);
        setInQueue(false);
        setSeconds(0);
        notifyWarning("Queue cancelled!");
    });
  };

  return  (
    <div className="mt-6 mx-auto w-2/5 flex justify-between">
      <Button 
        color="primary" 
        variant="solid" 
        className="text-lg py-5 mx-2" 
        fullWidth={true}
        onPress={startQueue}
        isLoading={inQueue}
        isDisabled={active}
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
