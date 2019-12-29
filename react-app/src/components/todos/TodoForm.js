import React, { Component } from 'react';
import { Field, reduxForm } from 'redux-form';

class TodoForm extends Component {
    renderField = ({ input, label, meta: { touched, error } }) => {
        return (
            <div className={`${touched && error ? 'text-danger' : ''} form-group`}>
                <input placeholder={`Enter a ${label} name`} {...input} autoComplete='off' className='form-control' />
                {touched && error && (
                    <span>{error}</span>
                )}
            </div>
        );
    };

    onSubmit = formValues => {
        this.props.onSubmit(formValues);
    };

    render() {
        const btnText = `${this.props.initialValues ? 'Edit' : 'Add'}`;
        return (
            <div className='col-md-12'>
                <form onSubmit={this.props.handleSubmit(this.onSubmit)}>
                    <div className='row'>
                        <div className='col-md-10'>
                            <Field name='task' component={this.renderField} label='Task' />
                        </div>
                        <div className='col-md-2'>
                            <button type='submit' className='btn btn-primary btn-block'>{btnText}</button>
                        </div>
                    </div>
                </form>
            </div>
        );
    }
}

const validate = formValues => {
    const errors = {};

    if (!formValues.task) {
        errors.task = 'Please enter at least 1 character';
    }

    return errors;
};

export default reduxForm({
    form: 'todoForm',
    touchOnBlur: false,
    validate
})(TodoForm);