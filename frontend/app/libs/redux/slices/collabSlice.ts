import {createSlice} from "@reduxjs/toolkit";
import {AppState} from "../store";


export interface CollabState {
    isLeaving : boolean
    isChatOpen : boolean
}

const initialState : CollabState = {
    isLeaving : false,
    isChatOpen : false,
};

export const collabSlice = createSlice({
    name: "collab",
    initialState,
    reducers: {
        setIsLeaving: (state, action) => {
            state.isLeaving = action.payload;
        },
        setIsChatOpen: (state, action) => {
            state.isChatOpen = action.payload;
        },
    },
});

export const { setIsLeaving, setIsChatOpen } = collabSlice.actions;
export const selectCollabState = (state : AppState) => state.collab;
export default collabSlice.reducer;
