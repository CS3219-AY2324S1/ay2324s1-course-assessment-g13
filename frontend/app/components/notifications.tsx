import { toast } from 'react-toastify';

export const notifyWarning = (warning : string) => toast.warn(warning, {
  theme: "dark"
});

export const notifySuccess = (message : string) => toast.success(message, {
  theme: "dark"
});

export const notifyError = (err : string) => toast.error(err, {
  theme:"dark"
});
