import axios from 'axios';
import { reset } from 'redux-form';
import { GET_TODOS, GET_TODO, ADD_TODO, DELETE_TODO, EDIT_TODO } from './types';
import history from '../history';

// GET TODOS
export const getTodos = () => async dispatch => {
    const res = await axios.get('/api/v1/product/');
    dispatch({
        type: GET_TODOS,
        payload: res.data
    });
};

export const addTodo = formValues => async dispatch => {
    const res = await axios.post('/api/v1/product/', { ...formValues });
    dispatch({
        type: ADD_TODO,
        payload: res.data
    });
    dispatch(reset('todoForm'));
};

export const getTodo = id => async dispatch => { // added
    const res = await axios.get(`/api/v1/product/${id}/`);
    dispatch({
        type: GET_TODO,
        payload: res.data
    });
};

export const deleteTodo = id => async dispatch => { // added
    await axios.delete(`/api/v1/product/${id}/`);
    dispatch({
        type: DELETE_TODO,
        payload: id
    });
    history.push('/');
};

export const editTodo = (id, formValues) => async dispatch => {
    const res = await axios.patch(`/api/v1/product/${id}/`, formValues);
    dispatch({
        type: EDIT_TODO,
        payload: res.data
    });
    history.push('/');
};