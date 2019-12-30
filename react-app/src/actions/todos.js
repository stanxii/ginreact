import axios from 'axios';
import { reset } from 'redux-form';
import { GET_TODOS, GET_TODO, ADD_TODO, DELETE_TODO, EDIT_TODO } from './types';
import history from '../history';


export const getTodos = () => async dispatch => {
    const response = await axios.get('http://localhost:8000/api/v1/todo/');
    dispatch({
        type: GET_TODOS,
        payload: response
    });
};

export const addTodo = formValues => async dispatch => {
    const response = await axios.post('http://localhost:8000/api/v1/todo/', { ...formValues });
    dispatch({
        type: ADD_TODO,
        payload: response
    });
    dispatch(reset('todoForm'));
};

export const getTodo = id => async dispatch => {
    const response = await axios.get(`http://localhost:8000/api/v1/todo/${id}/`);
    dispatch({
        type: GET_TODO,
        payload: response
    });
};

export const deleteTodo = id => async dispatch => {
    await axios.delete(`http://localhost:8000/api/v1/todo/${id}/`);
    dispatch({
        type: DELETE_TODO,
        payload: id
    });
    history.push('/');
};

export const editTodo = (id, formValues) => async dispatch => {
    const response = await axios.patch(`http://localhost:8000/api/v1/todo/${id}/`, formValues);
    dispatch({
        type: EDIT_TODO,
        payload: response
    });
    history.push('/');
};