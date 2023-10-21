"use client";

import { Divider } from "@nextui-org/divider";
import { Input } from "@nextui-org/input";
import { Select, SelectItem } from "@nextui-org/select";
import { LANGUAGES } from "../constants/languages";
import useUser from "../hooks/useUser";
import { useForm, Controller } from "react-hook-form";
import { Button } from "@nextui-org/button";
import { UserState, updateUser } from "../libs/redux/slices/userSlice";
import { Avatar } from "@nextui-org/avatar";
import ImageUpload from "../components/form/ImageUpload";
import useAuth from "../hooks/useAuth";
import { GET, PUT } from "../libs/axios/axios";
import { useDispatch } from "react-redux";
import { AxiosResponse } from "axios";
import { notifyError, notifySuccess } from "../components/toast/notifications";
import { LoginResponse, UserResponse } from "../(auth)/login/page";
import DeleteAccountModal from "../components/modal/deleteAccountModal";
import { update } from "../libs/redux/slices/authSlice";

interface UpdateUserRequest {
    username: string;
    photo_url: string;
    preferred_language: string;
}

export default function Profile() {
    const currentUserInfo = useUser();
    const { authId, role } = useAuth();
    const dispatch = useDispatch();

    const {
        control,
        handleSubmit,
        setValue,
        watch,
        reset,
        formState: { isDirty, errors },
      } = useForm({
        defaultValues: currentUserInfo
      });

    const onSubmit = handleSubmit(async (data: UserState) => {
        const requestBody: UpdateUserRequest = {
            username: data.username,
            photo_url: data.photoUrl,
            preferred_language: data.preferredLanguage
        }
        try {
            const response: AxiosResponse<UserResponse> = await PUT(`/users/${authId}`, requestBody);
            const {message, user} = response.data;
            dispatch(updateUser(user));
            notifySuccess(message);
            reset({}, {keepValues: true})
        } catch (error) {
            const message = error.message.data.message;
            notifyError(message);
        }
    })

    const handleChangeRole = async () => {
        try {
            const response: AxiosResponse<LoginResponse> = await GET(`/auth/user/${role === 'user' ? "upgrade" : "downgrade"}`);
            const { message, user } = response.data;
            dispatch(update(user))
            notifySuccess(message)
        } catch (error) {
            const message = error.message.data.message;
            notifyError(message);
        }
      }

    const photo = watch('photoUrl');
    
    const languages = LANGUAGES

    return (
        <>
            <div className="flex flex-col w-3/5 mx-auto mt-14">
                <div className="flex m-5 justify-between">
                    <div className="flex flex-col">
                        <h3 className="text-2xl font-bold">Profile Settings</h3>
                        <span>
                            You can update your profile here!
                        </span>
                    </div>
                    <Button color="primary" onClick={onSubmit} className=" my-1 " isDisabled={!isDirty}>
                        Save Changes
                    </Button>
                </div>
                <Divider />
                <div className="flex flex-row m-5 justify-between">
                    <p className="text-xl font-medium">Username</p>
                    <Controller 
                        rules={{required:{
                            value: true,
                            message: "Cannot be empty"
                        },
                        validate: {
                            notEmpty: (value) => value.trim() !== '' 
                            || 'Username cannot be empty or contain only whitespace',
                        }
                        }}
                        name="username"
                        control={control}
                        render={({field}) => (
                            <Input {...field} className="w-1/3" errorMessage={errors.username?.message as string}/>
                        )}
                    />
                </div>
                <Divider />
                <div className="flex flex-row m-5 justify-between">
                    <p className="text-xl font-medium">Preferred Language</p>
                    <Controller 
                        name="preferredLanguage"
                        control={control}
                        defaultValue={currentUserInfo.preferredLanguage}
                        render={({field: {onChange, value}}) => (
                            <>
                            <Select onChange={onChange} defaultSelectedKeys={[value]} className="w-1/3 h-1/2" labelPlacement="outside">
                                {languages.map((language) => (
                                    <SelectItem className="px-2" key={language} value={language}>
                                        {language}
                                    </SelectItem>
                                ))}
                            </Select>
                            </>
                        )}
                    />
                </div>
                <Divider />
                <div className="flex flex-row m-5 justify-between">
                    <p className="text-xl font-medium">Photo</p>
                    <div className="flex flex-col md:grid md:grid-cols-3 md:col-span-3">
                        <Avatar
                            showFallback
                            src={photo}
                            isBordered
                            color="primary"
                            className="h-20 w-20 self-center justify-self-center md: mb-4"
                        />
                        <div className="col-span-2 col-start-2">
                            <Controller
                                name="photoUrl"
                                control={control}
                                render={({ field }) => (
                                <ImageUpload
                                    setImage={value => {
                                    return setValue('photoUrl', value, { shouldDirty: true });
                                    }}
                                    {...field}
                                />
                                )}
                            />
                        </div>
                    </div>    
                </div>
                <Divider className="my-2" />
                <div className="flex flex-row m-5 justify-between">
                    <div className="flex flex-col">
                        <p className="text-2xl font-medium my-1">{role === 'admin' ? "Downgrade" : "Upgrade"} Account</p>
                        {
                            role === "admin" ?
                                (
                                    <p className="text-start text-sm my-1 md:text-base">
                                        Remove access to <span className="font-medium">add/delete</span> questions.
                                    </p>
                                ) :
                                (
                                    <p className="text-start text-sm my-1 md:text-base">
                                        Gain access to <span className="font-medium">add/delete</span> questions.
                                    </p>
                                )
                        }
                    </div>
                    <div className="flex flex-row items-center my-1">
                        <Button color="primary" onClick={handleChangeRole} className=" my-1">
                            {role === 'admin' ? "Downgrade" : "Upgrade"} Account
                        </Button>
                    </div>
                </div>
                <Divider className="my-2" />
                <div className="flex flex-row m-5 justify-between">
                    <div className="flex flex-col">
                        <p className="text-2xl font-medium my-1">Delete Account</p>
                        <p className="text-start text-sm my-1 md:text-base">
                            Deleting your account will remove all your information from our database.
                        </p>
                        <p className="text-start text-sm font-medium md:text-base">
                            This cannot be undone.
                        </p>
                    </div>
                    <div className="flex flex-row items-center my-1">
                        <DeleteAccountModal />
                    </div>
                </div>
            </div>
        </>
    )
}