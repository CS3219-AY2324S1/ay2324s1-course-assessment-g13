import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import { Question } from '../../types/question';

export interface QuestionBankState {
  questionBank: Question[];
}

const initialState: QuestionBankState = {
  questionBank: [],
};

export const questionBankSlice = createSlice({
  name: 'questionBank',
  initialState,
  reducers: {
    addQuestion: (state, action: PayloadAction<Question>) => {
      state.questionBank = [...state.questionBank, action.payload];
    },
  },
});

export const { addQuestion } = questionBankSlice.actions;

export default questionBankSlice.reducer;
