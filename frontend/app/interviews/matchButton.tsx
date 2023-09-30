import { Button } from "@nextui-org/react";
import { notifyError, notifyWarning } from "../components/notifications";
import { useEffect, useState } from "react";

export default function MatchButton({inQueue, setInQueue}) {
  const [seconds, setSeconds] = useState(0);
  const timeLimit = 30;

  useEffect(() => {
    if (seconds == timeLimit) {
      setInQueue(false);
      setSeconds(0);
      notifyError("Timeout! No suitable partner was found")
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
  }

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
