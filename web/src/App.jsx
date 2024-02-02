import './main.css';

const App = () => {
    return (
        <div id="app">
            <div id="app-header">
                <h1>Budget App</h1>
            </div>
            <div id="app-body">
                <p>Table or Flex Rows..?</p>
                <p>That is the question.</p>
            </div>
            <br />
            <h1>Table</h1>
            <div id="app-table">
                <table>
                    <tbody>
                        <tr class="flex my-4 p-8 bg-slate-800">
                            <td class="mx-8 px-8">Row 1, Column 1</td>
                            <td class="mx-8 px-8">Row 1, Column 2</td>
                            <td class="mx-8 px-8">Row 1, Column 3</td>
                        </tr>
                        <tr class="flex my-4 p-8 bg-slate-800">
                            <td class="mx-8 px-8">Row 2, Column 1</td>
                            <td class="mx-8 px-8">Row 2, Column 2</td>
                            <td class="mx-8 px-8">Row 2, Column 3</td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <br />
            <h1>Flex Rows</h1>
            <div id="app-flex" class="flex-row">
                <div class="flex my-4 p-8 bg-slate-800">
                    <div class="mx-8 px-8">Row 1, Column 1</div>
                    <div class="mx-8 px-8">Row 1, Column 2</div>
                    <div class="mx-8 px-8">Row 1, Column 3</div>
                </div>
                <div class="flex my-4 p-8 bg-slate-800">
                    <div class="mx-8 px-8">Row 2, Column 1</div>
                    <div class="mx-8 px-8">Row 2, Column 2</div>
                    <div class="mx-8 px-8">Row 2, Column 3</div>
                </div>
            </div>
            <div id="app-footer">
                <p>They look the same but table has 2 extra levels of indent.. 🤢</p>

                <p>Looks like I will be using 💪 Rows for this.! 😊</p>
            </div>
        </div>

    );
};

export default App;
