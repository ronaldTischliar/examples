async function fetchFromServer () {
    console.log('web component start');
    const response = await fetch('login.json');
    const payload = await response.json();
    console.log('web component',payload);
}

export {fetchFromServer}