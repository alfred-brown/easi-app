import React from 'react';
import { useParams } from 'react-router-dom';
import { Button } from '@trussworks/react-uswds';

import Footer from 'components/Footer';
import Header from 'components/Header';
import MainContent from 'components/MainContent';
import PageWrapper from 'components/PageWrapper';
import SystemProfileHealthCard from 'components/SystemProfileHealthCard';
import NotFound from 'views/NotFound';
import { mockSystemInfo } from 'views/Sandbox/mockSystemData';

const SystemProfileHealth = () => {
  const { systemId } = useParams<{ systemId: string }>();

  const systemInfo = mockSystemInfo.find(
    mockSystem => mockSystem.id === systemId
  );

  if (systemInfo === undefined) {
    return <NotFound />;
  }

  return (
    <PageWrapper>
      <Header />
      <MainContent>
        <div className="grid-container">
          <div className="desktop:grid-col-10">
            <h1>Sandbox</h1>
            <h1>System Profile</h1>
            <hr />
            <p>Overview your system&apos;s health</p>
            <SystemProfileHealthCard
              heading="Authority to Operate (ATO)"
              body={
                <>
                  <p>
                    The implementation of a Federal Government information
                    system requires a formal Government Authorization to Operate
                    (ATO) for infrastructure systems and/or all application
                    systems developed, hosted and/or maintained on behalf of CMS
                  </p>
                  <a href="/sandbox">Learn more about ATO</a>
                </>
              }
              footer={<Button type="button">View ATO Information</Button>}
              status={systemInfo.atoStatus}
              statusText={systemInfo.atoStatusText}
            />
            <SystemProfileHealthCard
              heading="Section 508"
              body={
                <>
                  <p>
                    Section 508 of the Rehabilitation Act requires federal
                    agencies to develop, procure, maintain and use information
                    and communications technology (ICT) that is accessible to
                    people with disabilities.
                  </p>
                  <a href="/sandbox">Learn more about 508 Compliance</a>
                </>
              }
              footer={
                <Button type="button">View 508 Compliance Information</Button>
              }
              status={systemInfo.section508Status}
              statusText={systemInfo.section508StatusText}
            />
            <SystemProfileHealthCard
              heading="Technical Review Board (TRB)"
              body={
                <>
                  <p>
                    The TRB provides oversight to ensure IT investments are
                    consistent with CMS’s IT strategy.
                  </p>
                  <a href="/sandbox">Learn more about the TRB</a>
                </>
              }
              footer={<Button type="button">View TRB Information</Button>}
              status={systemInfo.trbStatus}
              statusText={systemInfo.trbStatusText}
            />
          </div>
        </div>
      </MainContent>
      <Footer />
    </PageWrapper>
  );
};

export default SystemProfileHealth;