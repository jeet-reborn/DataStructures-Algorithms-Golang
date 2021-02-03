import Header from './components/Headers/Header';
import ImageDisplay from './components/ImageDisplay/ImageDisplay'
import Scores from './components/Scores/Scores'
import './App.css';

function App() {

  return (
    <div class="App">
      <Header />
      <div className="left">
        <ImageDisplay></ImageDisplay>
      </div>
      <div className="right">
        <Scores></Scores>
      </div>
    </div>
  );
}

export default App;
