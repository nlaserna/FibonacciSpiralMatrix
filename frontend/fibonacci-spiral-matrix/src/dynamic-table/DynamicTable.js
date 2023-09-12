import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';

const DynamicTable = ({ matrix }) => {
    if (!matrix || matrix.length === 0) {
        return <p>No data available.</p>;
    }

    return (
        <table className="table table-bordered mt-3">
            <tbody>
            {matrix.map((row, rowIndex) => (
                <tr key={rowIndex}>
                    {row.map((cell, columnIndex) => (
                        <td key={columnIndex}>{cell}</td>
                    ))}
                </tr>
            ))}
            </tbody>
        </table>
    );
};

export default DynamicTable;
