import {Button} from "@nextui-org/react";
import {notifySuccess, notifyWarning} from "../components/toast/notifications";
import {useSelector} from "react-redux";
import {selectPreferenceState} from "../libs/redux/slices/matchPreferenceSlice";
import {POST} from "../libs/axios/axios";
import {selectUsername} from "../libs/redux/slices/userSlice";

export default function MatchButton({inQueue, setInQueue, setSeconds, matchNotfound, setIsCancelled, setShouldNotifyCancelled}) {
  const preferenceState = useSelector(selectPreferenceState)
  const userState = useSelector(selectUsername);

  const startQueue = () => {
    setInQueue(true);
    setIsCancelled(false);
    setShouldNotifyCancelled(false);
    getMatch().then(r => {
      if (r.status == 200) {
        const payload = r.data;
        // If there is a match, notify success and redirect
        if (payload["match_user"] != "") {
          notifySuccess(`Matched with ${payload["match_user"]}`);
          setInQueue(false);
          setSeconds(0);
          // TODO perform redirection here based on payload redirect url
          console.log(payload["room_id"])
          const ws = new WebSocket(`ws://localhost:5005/ws/${payload["room_id"]}`)
          // TODO: router push to collaboration page
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

  const getMatch = async () => {
    return await POST("/match/find", {
      "username":`${userState}`,
      "match_criteria":`${preferenceState.toLowerCase()}`
    });
  };

  const cancelMatch = async () => {
    return await POST("/match/cancel", {
      "username": `${userState}`
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
