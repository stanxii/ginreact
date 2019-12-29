import axios from 'axios';
import { reset } from 'redux-form';
import { GET_TODOS, GET_TODO, ADD_TODO, DELETE_TODO, EDIT_TODO } from './types';
import history from '../history';


export const getTodos = () => async dispatch => {
    const res = await axios.get('http://localhost:8080/api/v1/todo/', {
        headers: {
            'Access-Control-Allow-Origin': '*',
        },
    });
    dispatch({
        type: GET_TODOS,
        payload: res.data
    });
};

export const addTodo = formValues => async dispatch => {
    const res = await axios.post('http://localhost:8080/api/v1/todo/', { ...formValues });
    dispatch({
        type: ADD_TODO,
        payload: res.data
    });
    dispatch(reset('todoForm'));
};

export const getTodo = id => async dispatch => {
    const res = await axios.get(`http://localhost:8080/api/v1/todo/${id}/`);
    dispatch({
        type: GET_TODO,
        payload: res.data
    });
};

export const deleteTodo = id => async dispatch => {
    await axios.delete(`/api/v1/todo/${id}/`);
    dispatch({
        type: DELETE_TODO,
        payload: id
    });
    history.push('/');
};

export const editTodo = (id, formValues) => async dispatch => {
    const res = await axios.patch(`/api/v1/todo/${id}/`, formValues);
    dispatch({
        type: EDIT_TODO,
        payload: res.data
    });
    history.push('/');
};