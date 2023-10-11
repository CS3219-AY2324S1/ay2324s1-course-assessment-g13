import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {Complexity} from "../../../types/question";
import {AppState} from "../store";

// User matching criteria
export interface PreferenceState {
    preference: string;
}

// User matching criteria initial state defaulted to easy
const initialState : PreferenceState = {
    preference: "Easy"
};

export const preferenceSlice = createSlice({
    name: "preference",
    initialState,
    reducers: {
        setPreference: (state, action : PayloadAction<Complexity>) => {
            state.preference = action.payload;
        },
    },
});

export const { setPreference } = preferenceSlice.actions;
export const selectPreferenceState = (state : AppState) => state.preference.preference;
export default preferenceSlice.reducer;
