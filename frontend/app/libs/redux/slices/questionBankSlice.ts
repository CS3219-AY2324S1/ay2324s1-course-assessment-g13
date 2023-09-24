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
    deleteQuestion: (state, action: PayloadAction<string>) => {
      state.questionBank = state.questionBank.filter(question => question.title !== action.payload);
    },
  },
});

export const { addQuestion, deleteQuestion } = questionBankSlice.actions;

export default questionBankSlice.reducer;
