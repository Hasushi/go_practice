async function fetchAttach() {
    try {
        const response = await fetch('http://localhost:8080/attach')
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        // ダウンロードするためのBlobを取得
        const data = await response.blob();
        // BlobをURLに変換
        const url = URL.createObjectURL(data);
        // a要素を作成してダウンロードをトリガー
        const a = document.createElement('a');
        a.href = url;
    } catch (error) {
        console.error('Error fetching attach endpoint:', error);
    }
}

async function fetchInline() {
    try {
        const response = await fetch('http://localhost:8080/inline');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.blob();
        console.log('Inline endpoint response:', data);
    } catch (error) {
        console.error('Error fetching inline endpoint:', error);
    }
}

const attachButton = document.getElementById('attachmentButton');
const inlineButton = document.getElementById('inlineButton');
attachButton.addEventListener('click', fetchAttach);
inlineButton.addEventListener('click', fetchInline);