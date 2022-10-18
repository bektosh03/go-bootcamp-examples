import React from 'react';
import './registration.css';

const Registration = () => {
	return (
		<div>
			<h1>Registration</h1>
			<form id='registration'>
				<div className='name-input'>
					<label htmlFor='name'>Name</label>
					<input id='name' type='text' />
				</div>
			</form>
		</div>
	);
};

export default Registration;