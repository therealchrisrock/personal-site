import { disableBodyScroll, enableBodyScroll, clearAllBodyScrollLocks } from 'body-scroll-lock';

// Inactivity time in seconds
const INACTIVITY_MS = 15 * 1000;
let inactivityTimeout: number;
// Element to show/hide
let isVisible = false;
// Start the inactivity timer on page load
export async function init() {
    const container = document.createElement('div')
    container.id = 'cloud9-container'
    container.style.zIndex = "50"
    container.style.opacity = '0'
    container.style.visibility = 'hidden';
    container.style.pointerEvents = 'none'
    container.style.transition = 'opacity 1s ease-in, visibility 0s';
    document.body.appendChild(container)
    const CLOUD9 = await import('./cloud9.js')
    CLOUD9.init(container)
    resetInactivityTimer();


    // Utility function for fading in/out
    function fadeElement(element: HTMLElement, fadeIn: boolean, duration = 500): void {
        element.style.transition = `opacity ${duration}ms, ${fadeIn ? 'visibility 0s' : 'visibility 0s ease-in 1s'}`;
        element.style.opacity = fadeIn ? "1" : "0";
        element.style.visibility = fadeIn ? "visible" : "hidden";
        element.style.pointerEvents = fadeIn ? "auto" : "none"; // Prevent interaction when hidden
        if (fadeIn) {
            disableBodyScroll(element)
        } else {
            enableBodyScroll(element)
        }
        isVisible = fadeIn
    }

// Main inactivity handler

    function resetInactivityTimer() {
        clearTimeout(inactivityTimeout);
        inactivityTimeout = window.setTimeout(() => {
            // Show the container after inactivity
            if (!isVisible) fadeElement(container, true);
        }, INACTIVITY_MS);
    }

// Hide the container and reset the timer on user interaction
    function handleUserInteraction() {
        if (isVisible) fadeElement(container, false);
        resetInactivityTimer(); // Restart timer after hiding
    }


/**
 * List of user engagement events.
 * Add or remove event types based on your requirements.
 */
const userEngagementEvents = [
    'click',
    'mousedown',
    'mouseup',
    'mousemove',
    'touchstart',
    'touchend',
    'touchmove',
    'keydown',
    'keyup',
    'scroll',
    'focus',
    'blur'
  ];
const resetEvents = [
    'click',
    'keyup',
    'keydown'
]
    function handleUserEngagement(event: Event): void {
        console.log(`User engaged with a "${event.type}" event.`);
        if (isVisible && resetEvents.includes(event.type)) fadeElement(container, false);
            resetInactivityTimer(); // Restart timer after hiding
        // Put your custom logic here.
        // E.g., trigger analytics, mark the session as active, etc.
    }
  
  // Attach listeners for each event in the list
  userEngagementEvents.forEach((eventName) => {
    window.addEventListener(eventName, handleUserEngagement, { passive: true });
  });
}
