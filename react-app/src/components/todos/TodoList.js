import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { getTodos, deleteTodo } from '../../actions/todos';

class TodoList extends Component {
    componentDidMount() {
        this.props.getTodos();
    }

    render() {
        return (
            <div className='ui relaxed divided list' style={{ marginTop: '2rem' }}>
                {this.props.todos.map(todo => (
                    <div className='item' key={todo.id}>
                        <div className='right floated content'> // added
                            <Link
                                to={`/delete/${todo.id}`}
                                className='small ui negative basic button'
                            >
                                Delete
                            </Link>
                            <Link to={`/edit/${todo.id}`} className='header'>
                                {todo.task}
                            </Link>
                        </div>
                        <i className='large calendar outline middle aligned icon' />
                        <div className='content'>
                            <a className='header'>{todo.task}</a>
                            <div className='description'>{todo.created_at}</div>
                        </div>
                    </div>
                ))}
            </div>
        );
    }
}

const mapStateToProps = state => ({
    todos: Object.values(state.todos)
});

/**
 * Connects this component to the store.
 */
export default connect(
    mapStateToProps,
    { getTodos, deleteTodo }
)(TodoList);