import './App.css';
import 'primereact/resources/themes/lara-light-indigo/theme.css';
import 'primereact/resources/primereact.min.css';
import FileUploadComponent from './components/FileUpload';
import DataTableComponent from './components/DataTable';

function App() {
    return (
        <div className="App">
            <FileUploadComponent/>
            <DataTableComponent/>
        </div>
    );
}

export default App;
