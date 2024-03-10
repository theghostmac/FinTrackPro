import financeImage from '../assets/personal-finance.png';
import { Link } from 'react-router-dom';
import './LandingPage.css';

const LandingPage = () => {
    return (
        <div className="landing-container">
            <header className="header">
                <h1>FinTrackPro</h1>
                <nav className="navigation">
                    <Link to="/register" className="nav-link">Register</Link>
                    <Link to="/about" className="nav-link">About</Link>
                </nav>
            </header>
            <main className="main-content">
                <section className="intro-section">
                    <h2>...Your Personal Finance Tracker</h2>
                    <p>Don't fall trap to Parkinson's Law:</p>
                    <blockquote>
                        "As your income increases your expenses will rise to meet or exceed that increase."
                    </blockquote>
                </section>
                <section className="feature-image">
                    <img src={financeImage} alt="Personal Finance" />
                </section>
            </main>
            <footer className="footer">
                <p>Built with Golang and React - TS</p>
                <div className="social-links">
                    <a href="https://github.com/theghostmac/FinTrackPro" target="_blank" rel="noopener noreferrer">GitHub</a>
                    <Link to="/about">About</Link>
                    <a href="https://x.com/ghostmac9" target="_blank" rel="noopener noreferrer">Twitter</a>
                </div>
            </footer>
        </div>
    );
};

export default  LandingPage;