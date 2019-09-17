/**
 * Function balance
 * @param {Element} container - DOM Element to search inside for elements to balance
 * @param {selector} balanceHandle - CSS selector applied to elements to be balanced
 *
 * This function will balance lines of text such that lines
 * stacked on top of each other will be of similar length.
 *
 * This prevents widow words on new lines inside of a responsive
 * design where content is dynamic and subject to change.
 *
 * A default step size tolerance of 10px can be overriden by
 * supplying a data attribute to the element to be balanced.
 */
export default function (container = document, balanceHandle = '.balance') {
	const els = container.querySelectorAll(balanceHandle);

	Array.from(els).forEach(el => {
		// Don't rebalance lines
		if (el.getAttribute('data-balanced')) return;

		// Allow for overriding default step size per element
		const stepSize = el.getAttribute('data-step-size') || 10;

		// Set element to default width
		el.style.maxWidth = 'inherit';

		// Capture initial height and width
		const height = el.offsetHeight;
		const width = el.offsetWidth;

		// Bug fix for trying to balance display none elements
		if (height === 0) {
			return;
		}

		// Set initial height and width to monitor during shrinking
		let nextHeight = height;
		let nextWidth = width;

		// Shrink until height changes
		while (nextHeight === height) {
			nextWidth -= stepSize;

			el.style.maxWidth = `${nextWidth}px`;
			nextHeight = el.offsetHeight;
		}

		// Add back previous step and set max width to remove widow caused by shrinking
		el.style.maxWidth = `${nextWidth + (stepSize * 3)}px`;

		// Mark as balanced
		el.setAttribute('data-balanced', true);
	});
}
