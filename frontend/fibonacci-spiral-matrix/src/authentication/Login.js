import React, { Component } from 'react';
import { Navigate } from 'react-router-dom';

class Login extends Component {
    constructor() {
        super();
        this.state = {
            username: '',
            password: '',
            message: '',
            redirectToHome: false,
        };
    }

    handleInputChange = (event) => {
        const { name, value } = event.target;
        this.setState({
            [name]: value,
        });
    }

    handleSubmit = async (event) => {
        event.preventDefault();
        const { username, password } = this.state;

        try {
            const response = await fetch('/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ username, password }),
            });

            if (response.ok) {
                this.setState({
                    message: 'Login successful',
                    redirectToHome: true,
                    isAuthenticated: true,
                });
                this.handleLogin();
            } else {
                const data = await response.json();
                this.setState({ message: data.message });
            }
        } catch (error) {
            console.error(error);
        }
    }

    render() {
        const { username, password, message, redirectToHome } = this.state;

        if (redirectToHome) {
            return <Navigate to="/" />;
        }

        return (
            <div className="container">
                <div className="row justify-content-center">
                    <div className="col-md-6">
                        <div className="card">
                            <div className="card-body">
                                <h1 className="card-title">Login</h1>
                                <form onSubmit={this.handleSubmit}>
                                    <div className="form-group">
                                        <label htmlFor="username">Username:</label>
                                        <input
                                            type="text"
                                            name="username"
                                            value={username}
                                            onChange={this.handleInputChange}
                                            className="form-control"
                                            required
                                        />
                                    </div>
                                    <div className="form-group">
                                        <label htmlFor="password">Password:</label>
                                        <input
                                            type="password"
                                            name="password"
                                            value={password}
                                            onChange={this.handleInputChange}
                                            className="form-control"
                                            required
                                        />
                                    </div>
                                    <button type="submit" className="btn btn-primary">Login</button>
                                </form>
                                <p className="mt-3 text-danger">{message}</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default Login;
