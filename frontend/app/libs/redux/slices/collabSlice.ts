import {createSlice} from "@reduxjs/toolkit";
import {AppState} from "../store";


export interface CollabState {
    isLeaving : boolean
}

const initialState : CollabState = {
    isLeaving : false
};

export const collabSlice = createSlice({
    name: "collab",
    initialState,
    reducers: {
        setIsLeaving: (state, action) => {
            state.isLeaving = action.payload;
        },
    },
});

export const { setIsLeaving } = collabSlice.actions;
export const selectCollabState = (state : AppState) => state.collab.isLeaving;
export default collabSlice.reducer;