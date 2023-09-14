import axios from 'axios';
import { Question } from './types/question';

axios.defaults.baseURL = 'http://localhost:8080/questions';
axios.defaults.headers.common['Content-Type'] = 'application/json';
axios.defaults.headers.common['Accept'] = 'application/json';

export const getQuestions = async () => {
    return await axios.get("");
}

export const createQuestion = async (data : Question) => {
    return await axios.post("", data);
}

export const deleteQuestion = async (id : string) => {
    return await axios.delete(`/${id}`);
}