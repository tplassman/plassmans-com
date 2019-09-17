import './polyfills';

// Components
// import Header from './components/header';

// Utilities
import { instantiate } from './utilities/components';

/* eslint-disable quote-props */
const classMap = {
	// 'header': Header,
};
/* eslint-enable quote-props */

// Event handler functions
function handleDOMConentLoaded() {
	// Call component constructors
	instantiate(classMap);
}

// Add event listeners
document.addEventListener('DOMContentLoaded', handleDOMConentLoaded);
