package com.epam.asset.tracking.service;

import static org.assertj.core.api.Assertions.assertThat;
import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;
import static org.mockito.Mockito.doReturn;
import static org.mockito.Mockito.mock;
import java.util.Optional;
import org.junit.Before;
import org.junit.Test;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import com.epam.asset.tracking.domain.BusinessProvider;
import com.epam.asset.tracking.repository.BusinessProviderRepository;

public class UserServiceImplTest {

  UserServiceImpl instance;

  @Before
  public void setup() {
    instance = new UserServiceImpl();
    instance.userRepository = mock(BusinessProviderRepository.class);

  }

  @Test
  public void happyPath() {
    BusinessProvider bp = new BusinessProvider();
    bp.setUsername("someUserName");
    bp.setEmail("an@email.com");
    bp.setPassword("pass");
    
    doReturn(Optional.of(bp)).when(instance.userRepository).findByUsername("someUserName");
    UserDetails user = instance.loadUserByUsername("someUserName");
    assertThat(user.getPassword(), is("pass"));

  }

  
  @Test(expected = UsernameNotFoundException.class)
  public void userNotFound() {

    doReturn(Optional.empty()).when(instance.userRepository).findByUsername("someUserName");
    
    instance.loadUserByUsername("someUserName");

  }
  
}
