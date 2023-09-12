import React, { Component, useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Register from './authentication/Register';
import Login from './authentication/Login';
import DynamicTable from "./dynamic-table/DynamicTable";

class App extends Component {
    constructor() {
        super();
        this.state = {
            isAuthenticated: false,
        };

    }

    handleLogin = () => {
        this.setState({ isAuthenticated: true });
    }

    render() {
        return (
            <Router>
                <div className="App">
                    <Routes>
                        <Route path="/register" element={<Register/>} />
                        <Route path="/login" element={<Login/>} />
                        <Route path="/" element={<DynamicSpiralMatrixTable/>} />
                    </Routes>
                </div>
            </Router>
        );
    }
}

function DynamicSpiralMatrixTable() {
    const [matrix, setMatrix] = useState([]);
    const [rows, setRows] = useState(0);
    const [columns, setColumns] = useState(0);

    const handleRowsChange = (e) => {
        setRows(e.target.value);
    };

    const handleColumnsChange = (e) => {
        setColumns(e.target.value);
    };

    const handleCalculate = () => {
        const apiUrl = `/?columns=${columns}&rows=${rows}`;

        fetch(apiUrl)
            .then((response) => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then((data) => {
                setMatrix(data["rows"]);
            })
            .catch((error) => {
                console.error('Error fetching data:', error);
                setMatrix([]);
            });
    };

    return (
        <div className="container mt-5">
            <div className="left">
                <h1>Fibonacci Spiral</h1>
                <p>Matrix Properties</p>
            </div>
            <div className="row justify-content-left">
                <div className="col-md-3">
                    <div className="mb-3">
                        <label className="form-label">Number of Rows:</label>
                        <input
                            type="number"
                            className="form-control"
                            value={rows}
                            onChange={handleRowsChange}
                            min="1"
                        />
                    </div>
                </div>
                <div className="col-md-3">
                    <div className="mb-3">
                        <label className="form-label">Number of Columns:</label>
                        <input
                            type="number"
                            className="form-control"
                            value={columns}
                            onChange={handleColumnsChange}
                            min="1"
                        />
                    </div>
                </div>
                <div className="col-md-2">
                    <button
                        className="btn btn-secondary"
                        onClick={handleCalculate}
                    >
                        Calculate
                    </button>
                </div>
            </div>
            <DynamicTable matrix={matrix} />
        </div>
    );
}

export default App;