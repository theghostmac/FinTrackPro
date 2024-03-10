// import React from 'react';
import reactLogo from '../assets/react.svg'
import viteLogo from '../../public/vite.svg'

const LandingPage = () => {
    return (
        <div>
            <div>
                <a href="https://vitejs.dev" target="_blank">
                    <img src={viteLogo} className="logo" alt="Vite logo"/>
                </a>
                <a href="https://react.dev" target="_blank">
                    <img src={reactLogo} className="logo react" alt="React logo"/>
                </a>
            </div>
            <h1>Welcome to FinTrackPro</h1>
            <div className="card">
                <p>Your personal finance tracker.</p>
                <button>Register</button>
            </div>
            <p className="read-the-docs">
                Click on the Vite and React logos to learn more about the technologies powering FinTrackPro.
            </p>
        </div>
    );

};

export default  LandingPage;