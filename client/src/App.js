import Input from './components/Input';
import CacheList from './components/CacheList';
import CacheItem from './components/CacheItem';

function App() {
  return (
    <div className="flex-col justify-center mt-12">
      <Input/>
      <CacheList/>
      <CacheItem/>
    </div>
  );
}

export default App;
