import './main.css';

const App = () => {
    return (
        <div id="app">
            <Header />
        </div>
    );
};

const Header = () => {
    return (
        <header>
            <h1 id="title">SVT Template</h1>
            <p id="subtitle">SolidJS + Vite + TailwindCSS</p>
        </header>
    );
}

export default App;
