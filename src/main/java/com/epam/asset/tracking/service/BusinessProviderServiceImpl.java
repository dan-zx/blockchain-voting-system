package com.epam.asset.tracking.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Component;

import com.epam.asset.tracking.domain.BusinessProvider;
import com.epam.asset.tracking.repository.BusinessProviderRepository;

import java.util.Optional;

@Component
public class BusinessProviderServiceImpl implements BusinessProviderService {

	@Autowired
	private BusinessProviderRepository repository;

	@Autowired
	BCryptPasswordEncoder bCryptPasswordEncoder;

	@Override
	public BusinessProvider save(BusinessProvider entity) {
		//encode password
		entity.setPassword(bCryptPasswordEncoder.encode(entity.getPassword()));
		return repository.save(entity);
	}

	@Override
	public Optional<BusinessProvider> findUserbyUsername(String name) {
		return repository.findByUsername(name);
	}

}
